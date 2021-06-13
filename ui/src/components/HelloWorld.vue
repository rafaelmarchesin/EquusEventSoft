<template>
  <v-container>
    <v-row class="text-center">
      <!-- Podemos colocar o logo feito pelo Herick -->
      <v-col cols="12">
        <v-img
          src="require('../assets/logo.jpeg')"
          class="my-3"
          contain
          height="200"
        />
      </v-col>

      <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb-3">Bem Vindo ao EventSoft</h1>

        <p class="subheading font-weight-regular">
          Com este aplicativo, você conseguirá criar eventos,
          <br />cadastrar novos participantes e muito mais.
          <br />
          <router-link to="/team">Conheça Nossa Equipe</router-link>
        </p>
      </v-col>

      <v-col class="mb-5" cols="12">
        <h2 class="headline font-weight-bold mb-3">Próximos passos:</h2>

        <v-row justify="center">
          <router-link class="subheading mx-3" to="/events"
            >Criar um evento</router-link
          >
          <router-link class="subheading mx-3" to="/participants"
            >Cadastrar um participante</router-link
          >
          <router-link class="subheading mx-3" to="/about"
            >Tirar dúvidas</router-link
          >
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
  export default {
    name: 'HelloWorld',

    data: () => ({
      ecosystem: [
        {
          text: 'vuetify-loader',
          href: 'https://github.com/vuetifyjs/vuetify-loader',
        },
        {
          text: 'github',
          href: 'https://github.com/vuetifyjs/vuetify',
        },
        {
          text: 'awesome-vuetify',
          href: 'https://github.com/vuetifyjs/awesome-vuetify',
        },
      ],
      importantLinks: [
        {
          text: 'Documentation',
          href: 'https://vuetifyjs.com',
        },
        {
          text: 'Chat',
          href: 'https://community.vuetifyjs.com',
        },
        {
          text: 'Made with Vuetify',
          href: 'https://madewithvuejs.com/vuetify',
        },
        {
          text: 'Twitter',
          href: 'https://twitter.com/vuetifyjs',
        },
        {
          text: 'Articles',
          href: 'https://medium.com/vuetify',
        },
      ],
      whatsNext: [
        {
          text: 'Explore components',
          href: 'https://vuetifyjs.com/components/api-explorer',
        },
        {
          text: 'Select a layout',
          href: 'https://vuetifyjs.com/getting-started/pre-made-layouts',
        },
        {
          text: 'Frequently Asked Questions',
          href: 'https://vuetifyjs.com/getting-started/frequently-asked-questions',
        },
      ],
    }),
  }
</script>
