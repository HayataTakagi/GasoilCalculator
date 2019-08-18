/* eslint-disable */
<template>
    <v-app id="inspire">
        <v-app-bar app color="indigo" dark >
            <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
            <v-toolbar-title>GasOilGenerator</v-toolbar-title>
        </v-app-bar>

        <v-content class="pt-10">
            <v-container fill-height fluid align-start justify-center>
                <v-col cols="12">
                    <FeeForm :oilPrices="prices" />
                </v-col>
            </v-container>
        </v-content>
        <v-footer
                color="indigo"
                app
        >
            <span class="white--text">&copy; 2019</span>
        </v-footer>
    </v-app>
</template>

<script>
    import FeeForm from "../modules/FeeForm";
    import ResultDialog from "../modules/ResultDialog";
    import axios from "axios"

    export default {
        created () {
            const self = this;
            axios.get("http://localhost:1991/gasoil")
                .then((res) => {
                    console.log(res.data);
                    self.prices = res.data.fuels
                })
        },
        components: {
            FeeForm,
            ResultDialog,
        },
        props: {
            source: String,
        },
        data: () => ({
            drawer: null,
            article: {
                'title': 'ArticleTest',
                'body': 'ArticleBodyTest',
            },
            prices: {}
        }),
    }
</script>