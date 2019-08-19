<template>
    <v-dialog v-model="dialog" width="600px">
        <template v-slot:activator="{ on }">
            <v-btn color="primary" v-on="on" class="mr-4" :disabled="!valid" v-on:click.stop="$emit('validate')">計算する</v-btn>
        </template>
        <v-card>
            <v-card-title>
                <span class="headline">計算結果</span>
            </v-card-title>
            <v-card-text>ガソリン代: 約{{ parseInt(this.oilPrice).toLocaleString() }}円</v-card-text>
            <v-card-text>CO2排出量: 約{{ parseInt(this.co2Amount.co2).toLocaleString() }}kg</v-card-text>
            <v-card-text>杉の木: {{ parseFloat(this.co2Amount.tree).toFixed(1).toLocaleString() }}本</v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="green darken-1" text @click="dialog = false">閉じる</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
    export default {
        name: "ResultDialog",
        // data () {
        //     return {
        //         dialog: false,
        //     }
        // },
        props: {
            dialog: Boolean,
            oilPrice: Number,
            co2Amount: Object,
            valid: Boolean,
        },
    }
</script>

<style scoped>

</style>