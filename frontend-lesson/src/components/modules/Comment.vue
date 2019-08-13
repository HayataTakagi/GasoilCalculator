<template>
    <v-card
            class="mx-auto"
            color="#26c6da"
            dark
            max-width="400"
    >
        <v-card-title>
            <v-icon
                    large
                    left
            >
                mdi-comment
            </v-icon>
            <span class="title font-weight-light">Comment</span>
        </v-card-title>

        <v-card-text class="headline font-weight-bold">
            ここにコメントが表示されます
        </v-card-text>
        <v-card-actions>
            <v-list-item class="grow">
                <v-list-item-avatar color="grey darken-3">
                    <v-img
                            class="elevation-6"
                            src="https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortCurly&accessoriesType=Prescription02&hairColor=Black&facialHairType=Blank&clotheType=Hoodie&clotheColor=White&eyeType=Default&eyebrowType=DefaultNatural&mouthType=Default&skinColor=Light"
                    ></v-img>
                </v-list-item-avatar>

                <v-list-item-content>
                    <v-list-item-title>Evan You</v-list-item-title>
                </v-list-item-content>
            </v-list-item>
        </v-card-actions>
    </v-card>
</template>
