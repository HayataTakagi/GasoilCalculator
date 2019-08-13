<template>
    <v-app id="inspire">
        <v-navigation-drawer v-model="drawer"  app >
            <v-list dense>
                <v-list-item @click="">
                    <v-list-item-action>
                        <v-icon>article</v-icon>
                    </v-list-item-action>
                    <v-list-item-content>
                        <v-list-item-title>Article</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
                <v-list-item @click="">
                    <v-list-item-action>
                        <v-icon>comment</v-icon>
                    </v-list-item-action>
                    <v-list-item-content>
                        <v-list-item-title>Comment</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
            </v-list>
        </v-navigation-drawer>

        <v-app-bar app color="indigo" dark >
            <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
            <v-toolbar-title>Application</v-toolbar-title>
        </v-app-bar>

        <v-content>
            <v-container class="fill-height" fluid>
                <v-row align="center" justify="center" >
                    <Article :title="article.title" :body="article.body"/>
                    <Comment/>
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
    import Article from "../modules/Article";
    import Comment from "../modules/Comment";

    export default {
        components: {
            Article,
            Comment
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