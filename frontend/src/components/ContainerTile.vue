<template>
  <v-card>
    <v-card-title class="headline" primary-title>
      <div>{{container.Image}}</div>
      <v-spacer></v-spacer>
      <v-chip>
        <v-chip color="green" text-color="white" class="subtitle-1 ml-n3 mr-1 px-2">{{container.State}}</v-chip>
        <div class="caption">{{container.Status}}</div>
      </v-chip>
      <v-btn @click="logsView=true" :disabled="logsView" color="primary" text>
        <v-icon>mdi-code-equal</v-icon>
      </v-btn>  
      <v-btn @click="logsView=false" :disabled="!logsView" color="primary" text>
        <v-icon>mdi-code-braces</v-icon>
      </v-btn>
    </v-card-title>
    <v-card-text v-if="logsView" class="tile-text" v-bind:id="textID">
      <code v-bind:class="textID">{{logText}}</code>
    </v-card-text>
    <v-content v-if="!logsView">
      {{container}}
    </v-content>
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
      snackbarTimeout: 750,
      logsView: true
    };
  },
  mounted() {
    this.textID = "log-text-" + this.container.Id;
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
          if (scrollEl) {
            scrollEl.scrollTop = scrollEl.scrollHeight
          }
        }
        this.goToBottom()
      }, 50)
    }
  }
};
</script>

<style scoped>
</style>