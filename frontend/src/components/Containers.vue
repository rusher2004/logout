<template>
  <v-container>
    <v-layout>
      <v-flex xs12 sm6 offset-sm3>
        <div v-if="containers && containers.length">
          <div v-for="c in containers" v-bind:key="c.Hostname">
            <v-btn color="secondary" flat @click="chooseMe(c, $event)">{{c.Image}} - {{c.Hostname}}</v-btn>
          </div>
        </div>
      </v-flex>
    </v-layout>
    <div v-if="selectedContainer" class="text-xs-center">
      <v-dialog v-model="dialog" scrollable width="900">
        <v-card>
          <v-card-title class="headline" primary-title>{{selectedContainer.Image}}</v-card-title>
          <v-card-text id="log-text"><code>{{logText}}</code></v-card-text>
          <v-divider></v-divider>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" flat @click="dialog = false">Close</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      containers: [],
      dialog: false,
      selectedContainer: null,
      logText: ''
    };
  },
  created() {
    this.getContainers()
  },
  methods: {
    getContainers: function() {
      window.backend.GetContainers().then(result => {
        let formatted = result.map(x => JSON.parse(x))
        this.containers = formatted;
        console.log(formatted)
      });
    },
    chooseMe: async function(container, event) {
      this.selectedContainer = container
      this.dialog = true

      while (this.dialog) {
        let res = await window.backend.ContainerLogs(container.Hostname)
        // console.log(this.selectedContainer.Image)
        this.logText= res
        // console.log(this.$el)
        let textEl = document.getElementById("log-text")
        textEl.scrollTop = textEl.scrollHeight
        this.wait(100)
      }
    },
    wait: function (ms){
      var start = new Date().getTime();
      var end = start;
      while(end < start + ms) {
        end = new Date().getTime();
      }
    }
  }
};
</script>

<style scoped>
  #log-text {
    overflow: scroll;
  }
</style>