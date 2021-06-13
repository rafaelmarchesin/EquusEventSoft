<template>
  <v-container>
    <v-list-item>
      <v-list-item-content>
        <v-list-item-title class="text-h6"> Menu </v-list-item-title>
        <v-list-item-subtitle> Escolha uma opção abaixo </v-list-item-subtitle>
      </v-list-item-content>
    </v-list-item>

    <v-divider></v-divider>

    <v-list dense nav>
      <v-list-item link to="/">
        <v-list-item-icon>
          <v-icon>mdi-view-dashboard</v-icon>
        </v-list-item-icon>

        <v-list-item-content>
          <v-list-item-title>Página Inicial</v-list-item-title>
        </v-list-item-content>
      </v-list-item>

      <v-list-item link to="/events">
        <v-list-item-icon>
          <v-icon>mdi-view-dashboard</v-icon>
        </v-list-item-icon>

        <v-list-item-content>
          <v-list-item-title>Eventos</v-list-item-title>
        </v-list-item-content>
      </v-list-item>

      <v-list-item link to="/participants">
        <v-list-item-icon>
          <v-icon>mdi-view-dashboard</v-icon>
        </v-list-item-icon>

        <v-list-item-content>
          <v-list-item-title>Participantes</v-list-item-title>
        </v-list-item-content>
      </v-list-item>

      <v-list-item link to="/reports">
        <v-list-item-icon>
          <v-icon>mdi-view-dashboard</v-icon>
        </v-list-item-icon>

        <v-list-item-content>
          <v-list-item-title>Relatórios</v-list-item-title>
        </v-list-item-content>
      </v-list-item>

      <v-list-item link to="/team">
        <v-list-item-icon>
          <v-icon>mdi-view-dashboard</v-icon>
        </v-list-item-icon>

        <v-list-item-content>
          <v-list-item-title>Time do EventSoft</v-list-item-title>
        </v-list-item-content>
      </v-list-item>

      <v-list-item link to="/about">
        <v-list-item-icon>
          <v-icon>mdi-view-dashboard</v-icon>
        </v-list-item-icon>

        <v-list-item-content>
          <v-list-item-title>Sobre</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list>
  </v-container>
</template>
