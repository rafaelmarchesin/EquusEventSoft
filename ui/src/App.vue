<template>
  <v-app>
    <v-navigation-drawer v-model="drawer" app class="white">
      <Menu />
    </v-navigation-drawer>

    <v-app-bar app color="primary" dark>
      <div class="d-flex align-center">
        <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
        <v-toolbar-title>EventSoft</v-toolbar-title>
      </div>

      <v-spacer></v-spacer>

      <v-btn text to="/login">
        <span class="mr-2">Login</span>
      </v-btn>
    </v-app-bar>

    <v-main>
      <router-view />
    </v-main>
    <v-main>
    </v-main>
  </v-app>
</template>

<script>
import HelloWorld from './components/HelloWorld';
import Menu from './components/Menu.vue'

export default {
  name: 'App',

  components: {
    HelloWorld,
	Menu,
  },

  data: () => ({
    drawer: true,
  }),
};
</script>
