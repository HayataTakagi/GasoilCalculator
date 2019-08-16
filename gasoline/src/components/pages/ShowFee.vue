/* eslint-disable */
<template>
    <v-app id="inspire">
        <v-app-bar app color="indigo" dark >
            <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
            <v-toolbar-title>Application</v-toolbar-title>
        </v-app-bar>

        <v-content>
            <v-container fill-height fluid>
                <v-row align="start" justify="center" >
                    <FeeForm/>
                </v-row>
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

    export default {
        components: {
            FeeForm,
            ResultDialog
        },
        props: {
            source: String,
        },
        data: () => ({
            drawer: null,
            article: {
                'title': 'ArticleTest',
                'body': 'ArticleBodyTest',
            }
        }),
    }
</script>