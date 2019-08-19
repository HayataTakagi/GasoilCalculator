<template>
    <v-form
            ref="form"
            v-model="valid"
            lazy-validation
    >
        <v-radio-group v-model="oil_type" :mandatory="true" label="油種">
            <v-radio :label="`レギュラー ${oilPrices.regular.price}(¥/ℓ)`" value="regular"></v-radio>
            <v-radio :label="`ハイオク ${oilPrices.high_octane.price}(¥/ℓ)`" value="high_octane"></v-radio>
            <v-radio :label="`軽油 ${oilPrices.light.price}(¥/ℓ)`" value="light"></v-radio>
            <v-radio label="任意" value="optional"></v-radio>
        </v-radio-group>
        <v-text-field
                v-model="userOilPrice"
                :rules="[v => !isNaN(parseInt(v, 10)) || '数値を入力してください',]"
                label="ガソリン価格"
                :disabled="!isUserOilPrice"
        ></v-text-field>

        <v-text-field
                v-model="roadLength"
                :rules="lengthRules"
                label="走行距離"
                required
        ></v-text-field>

        <v-autocomplete
                v-model="selectCarFuel"
                :items="cars"
                :loading="isLoading"
                :search-input.sync="search"
                :hint="`${Object.keys(selectCarFuel).length ? '燃費: ' + selectCarFuel.fuelEconomy + ' (km/l)' : '車種を選択して下さい'}`"
                color="primary"
                hide-no-data
                hide-selected
                item-text="name"
                item-value="fuelEconomy"
                label="車種"
                placeholder="Start typing to Search"
                prepend-icon="mdi-database-search"
                persistent-hint
                return-object
                :rules="[v => Object.keys(selectCarFuel).length !== 0 || '車種を選択して下さい']"
                required
        ></v-autocomplete>

        <ResultDialog :oilPrice="calculatePrice()" :co2Amount="calculateCo2()" :valid="valid" @validate="validate" :dialog="dialog"/>
        <v-btn
                color="error"
                class="mr-4"
                @click="reset"
        >
            Reset Form
        </v-btn>

    </v-form>
</template>

<script>
    import ResultDialog from "./ResultDialog";

    export default {
        components: {
            ResultDialog,
        },
        name: "FeeForm",
        data: () => ({
            valid: false,
            roadLength: 150,
            userOilPrice: 120,
            lengthRules: [
                v => !!v || '距離は必須項目です',
                v => !isNaN(parseInt(v, 10)) || '数値を入力してください',
            ],
            checkbox: false,
            oil_type: "regular",
            selectCarFuel: {},
            descriptionLimit: 60,
            cars: [],
            isLoading: false,
            model: null,
            search: null,
            dialog: false
        }),
        methods: {
            calculatePrice: function() {
                const result = ( this.roadLength / this.selectCarFuel.fuelEconomy ) * this.oilPrice;
                return result;
            },
            calculateCo2: function() {
                const CO2_MATRIX = {
                    regular: 2.3,
                    high_octane: 2.3,
                    light: 2.6,
                    optional: 2.3,
                };
                const TREE_CO2 = 14; // 一年に杉の木が吸収するCO2(kg)
                const amountFuel = this.roadLength / this.selectCarFuel.fuelEconomy;
                const co2Amount = amountFuel * CO2_MATRIX[this.oil_type];
                const treeAmount = (co2Amount / TREE_CO2).toFixed(2);
                return {co2: co2Amount, tree: treeAmount};
            },
            reset () {
                this.oil_type = "regular";
                this.roadLength = 0;
                this.selectCarFuel = {};
            },
            validate () {
                if(this.$refs.form.validate()) {
                    this.dialog = false;
                }
            },
        },
        props: {
            oilPrices: Object,
        },
        computed: {
            isUserOilPrice: function () {
                return this.oil_type === "optional";
            },
            oilPrice: function () {
                return this.isUserOilPrice ? this.userOilPrice : this.oilPrices[this.oil_type].price;
            },
        },
        watch: {
            search (val) {
                // Items have already been loaded
                if (this.cars.length > 0) return;

                // Items have already been requested
                if (this.isLoading) return;

                this.isLoading = true;

                // Lazily load input items
                fetch('http://localhost:1991/car')
                    .then(res => res.json())
                    .then(res => {
                        const { count, cars } = res;
                        this.count = count;
                        this.cars = cars
                    })
                    .catch(err => {
                        console.log(err)
                    })
                    .finally(() => (this.isLoading = false))
            },
        },
    }
</script>

<style scoped>

</style>