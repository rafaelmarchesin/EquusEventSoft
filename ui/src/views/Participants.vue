<template>
  <v-container>
    <h1>Participantes</h1>
    <br />
    <v-form class="px-3" ref="form">
      <v-text-field
        v-model="Name"
        counter
        maxlength="250"
        label="Nome do Participante"
      ></v-text-field>

      <v-text-field
        v-model="Age"
        counter
        label="Idade"
      ></v-text-field>

      <v-text-field
        v-model="Email"
        counter
        maxlength="250"
        type="mail"
        hint="email@provedor.com"
        label="E-mail"
      ></v-text-field>

      <v-text-field
        v-model="Course"
        counter
        label="Curso"
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
      Name: "",
      Email: "",
      Age: "",
      Course: "",
    };
  },

  methods: {
    createEvent() {
      this.$refs.form;
      console.log(this.eventName);
    },

    async gcall() {
      console.log(`Nome do Evento: ${this.eventName}`);

      const data = {
        Name: this.Name,
        Email: this.Email,
        Age: this.Age,
        Course: this.Course,
      };
      // You can await too.
      var create_partcipant_data = await g.call("create_events", data);

      console.log(`${create_partcipant_data}`);
    },
  },
});
</script>
