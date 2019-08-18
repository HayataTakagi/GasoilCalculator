<template>
    <v-form
            ref="form"
            v-model="valid"
            lazy-validation
    >
        <v-radio-group v-model="oil_type" :mandatory="true">
            <v-radio :label="`レギュラー ${oilPrices.regular.price}(¥/ℓ)`" value="regular"></v-radio>
            <v-radio :label="`ハイオク ${oilPrices.high_octane.price}(¥/ℓ)`" value="high_octane"></v-radio>
            <v-radio :label="`軽油 ${oilPrices.light.price}(¥/ℓ)`" value="light"></v-radio>
        </v-radio-group>

        <v-text-field
                v-model="roadLength"
                :rules="lengthRules"
                label="走行距離"
                required
        ></v-text-field>

        <v-select
                v-model="selectFuel"
                :items="items"
                item-text="carName"
                item-value="fuelEconomy"
                :hint="`燃費: ${selectFuel.fuelEconomy} (km/l)`"
                :rules="[v => !!v || 'Item is required']"
                label="車種"
                persistent-hint
                return-object
                required
        ></v-select>

        <ResultDialog :oilPrice="calculate()" :valid="valid"/>
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
            valid: true,
            roadLength: 150,
            lengthRules: [
                v => !!v || '距離は必須項目です',
                v => !isNaN(parseInt(v, 10)) || '数値を入力してください',
            ],
            selectFuel: { carName: 'ハイブリッド車 [20(km/ℓ)]', fuelEconomy: 20 },
            items: [
                { carName: 'ハイブリッド車 [20(km/ℓ)]', fuelEconomy: 20 },
                { carName: 'ガソリン車 [10(km/ℓ)]', fuelEconomy: 10 },
                { carName: '大型車 [8(km/ℓ)]', fuelEconomy: 8 },
            ],
            checkbox: false,
            oil_type: "regular"
        }),
        methods: {
            calculate: function() {
                const result = ( this.roadLength / this.selectFuel.fuelEconomy ) * this.oilPrice;
                return result;
            },
            reset () {
                this.oil_type = "regular";
                this.oilPrice = this.oilPrices.regular.price;
                this.roadLength = 0;
                this.selectFuel = { carName: 'ハイブリッド車 [20(km/ℓ)]', fuelEconomy: 20 };
                this.valid = false;
            },
        },
        props: {
            oilPrices: Object,
        },
        computed: {
            oilPrice: function () {
                return this.oilPrices[this.oil_type].price
            },
        },
    }
</script>

<style scoped>

</style>