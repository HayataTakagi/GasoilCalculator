
<template>
    <div>
        <v-card
                class="mx-auto"
                :flat="flat"
                :loading="loading"
                :outlined="outlined"
                :elevation="elevation"
                :raised="raised"
                :width="width"
                :height="height"
        >
            <v-img
                    v-if="media"
                    class="white--text"
                    height="200px"
                    src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
            >
                <v-card-title class="align-end fill-height">{{ this.title }}</v-card-title>
            </v-img>
            <v-card-title v-else>2019/08/14 </v-card-title>
            <v-card-text>{{ this.body }}</v-card-text>
        </v-card>
    </div>
</template>

<script>
    export default {
        data: () => ({
            flat: false,
            media: true,
            loading: false,
            actions: true,
            outlined: false,
            elevation: undefined,
            raised: false,
            width: 800,
            height: undefined,
        }),
        props: {
            title: String,
            body: String
        },
    }
</script>

<style scoped>
    .v-input__slider {
        width: 100%;
    }
</style>