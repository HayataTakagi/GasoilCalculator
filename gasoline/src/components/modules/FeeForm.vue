<template>
    <v-form
            ref="form"
            v-model="valid"
            lazy-validation
    >
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
            ResultDialog
        },
        name: "FeeForm",
        data: () => ({
            oilPrice: 120,
            valid: true,
            roadLength: 150,
            lengthRules: [
                v => !!v || '距離は必須項目です',
                v => !isNaN(parseInt(v, 10)) || '数値を入力してください',
            ],
            selectFuel: { carName: 'ハイブリッド車', fuelEconomy: 20 },
            items: [
                { carName: 'ハイブリッド車', fuelEconomy: 20 },
                { carName: 'ガソリン車', fuelEconomy: 10 },
                { carName: '大型車', fuelEconomy: 8 },
            ],
            checkbox: false,
        }),
        methods: {
            calculate: function() {
                const result = ( this.roadLength / this.selectFuel.fuelEconomy ) * this.oilPrice;
                return result;
            },
            reset () {
                this.roadLength = "";
                this.selectFuel = { carName: 'ハイブリッド車', fuelEconomy: 20 };
                this.valid = false;
            },
        }
    }
</script>

<style scoped>

</style>