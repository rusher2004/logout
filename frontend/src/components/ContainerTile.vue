<template>
  <v-card>
    <v-card-title class="headline" primary-title>
      <div>{{container.Image}}</div>
      <v-spacer></v-spacer>
      <v-btn color="primary" text>
        <v-icon>mdi-code-braces</v-icon>
      </v-btn>
    </v-card-title>
    <v-card-text class="tile-text" v-bind:id="textID">
      <code v-bind:class="textID">{{logText}}</code>
    </v-card-text>
    <v-divider></v-divider>
    <v-card-actions>
      <v-switch class="ml-1" v-model="follow" inset :label="'Follow'"></v-switch>
      <v-spacer></v-spacer>
      <div>
        <v-btn color="primary" text @click="copy">
          <v-icon>mdi-content-copy</v-icon>
        </v-btn>
        <v-btn color="primary" text @click="dialog = false">
          <v-icon>mdi-arrow-expand</v-icon>
        </v-btn>
        <v-snackbar
          v-model="snackbar"
          :timeout="snackbarTimeout"
          :top="true"
          :right="true"
          :color="'success'"
          >
          Logs copied to clipboard!
        </v-snackbar>
      </div>
    </v-card-actions>
  </v-card>
</template>

<script>
import Wails from "@wailsapp/runtime";

export default {
  props: {
    container: Object
  },
  data() {
    return {
      logText: "",
      textID: "",
      follow: true,
      snackbar: false,
      snackbarTimeout: 500
    };
  },
  created() {
    this.textID = "log-text-" + this.container.Id;
  },
  mounted() {
    window.backend.Containers.StreamContainerLogs(this.container.Id)
    Wails.Events.On(`log-stream-${this.container.Id}`, message => {
      let arr = message.log.substring(1, message.log.length - 1).split(" ");
      let text = String.fromCharCode.apply(null, arr);
      this.logText += text;
    });
    this.goToBottom()
  },
  methods: {
    copy: function() {
      navigator.clipboard.writeText(this.logText)
      this.snackbar = true
    },
    goToBottom: function() {
      setTimeout(() => {
        if (this.follow) {
          let scrollEl = document.getElementById(this.textID)
          scrollEl.scrollTop = scrollEl.scrollHeight
        }
        this.goToBottom()
      }, 50)
    }
  }
};
</script>

<style scoped>
</style>