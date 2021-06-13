<template>
  <v-container>
    <h1>Eventos</h1>
    <br />
    <v-form class="px-3" ref="form">
      <v-text-field
        v-model="eventName"
        counter
        maxlength="250"
        label="Nome do Evento"
      ></v-text-field>

      <v-text-field
        v-model="eventeResponsible"
        counter
        maxlength="250"
        label="Nome do Responsável"
      ></v-text-field>

      <v-text-field
        v-model="responsibleEmail"
        counter
        maxlength="250"
        type="mail"
        hint="email@provedor.com"
        label="E-mail do Responsável"
      ></v-text-field>

      <v-menu
        ref="menu"
        v-model="name"
        :close-on-content-click="false"
        :return-value.sync="date"
        transition="scale-transition"
        offset-y
        min-width="auto"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
            v-model="eventDate"
            label="Escolha uma data"
            append-icon="mdi-calendar"
            hint="DD/MM/YYYY format"
            readonly
            v-bind="attrs"
            v-on="on"
          ></v-text-field>
        </template>
        <v-date-picker v-model="date" no-title scrollable>
          <v-spacer></v-spacer>
          <v-btn text color="primary" @click="menu = false"> Cancelar </v-btn>
          <v-btn text color="primary" @click="$refs.menu.save(date)">
            OK
          </v-btn>
        </v-date-picker>
      </v-menu>

      <v-text-field
        v-model="eventTime"
        label="Horário"
        value="12:30:00"
        type="time"
        suffix="BRT"
      ></v-text-field>
      <v-text-field
        v-model="eventLocation"
        counter
        maxlength="250"
        hint="This field uses maxlength attribute"
        label="Local de Realização"
      ></v-text-field>
      <br />
      <v-btn block @click="gcall()"> Cadastrar </v-btn>
    </v-form>
  </v-container>
</template>

<script>
import Vue from "vue";
import g from "guark";

export default Vue.extend({
  name: "Events",

  data: function () {
    return {
      eventName: '',
      eventeResponsible: '',
      responsibleEmail: '',
      eventDate: null,
      eventTime: null,
      eventLocation: '',
    }
  },

  methods: {
    createEvent() {
      this.$refs.form
      console.log(this.eventName);
    },

    async gcall() {
      console.log(`Nome do Evento: ${this.eventName}`)

      const data = {
        eventName: this.eventName,
        eventeResponsible: this.eventeResponsible,
        responsibleEmail: this.responsibleEmail,
        eventDate: this.eventDate,
        eventTime: this.eventTime,
        eventLocation: this.eventLocation,
      }
      // You can await too.
      var create_events_data = await g.call("create_events", data);

      console.log(`${create_events_data}`);
    },
  },
});
</script>
