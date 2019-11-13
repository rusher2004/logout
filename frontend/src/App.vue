<template>
  <v-app id="inspire">
    <!-- <v-navigation-drawer v-model="drawer" app clipped>
      <v-list dense>
        <v-list-item>
          <v-list-item-content>
            <div class="d-flex align-center">
              <v-icon>mdi-brightness-7</v-icon>
              <v-switch class="ml-3" v-model="$vuetify.theme.dark" inset></v-switch>
              <v-icon>mdi-brightness-2</v-icon>
            </div>
          </v-list-item-content>
        </v-list-item>
        <v-list-item link>
          <v-list-item-action>
            <v-icon>mdi-view-dashboard</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>Dashboard</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item link>
          <v-list-item-action>
            <v-icon>mdi-settings</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>Settings</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer> -->

    <v-app-bar app clipped-left>
      <!-- <v-app-bar-nav-icon @click.stop="drawer = !drawer" /> -->
      <v-toolbar-title>logOut</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items>
        <v-icon>mdi-brightness-7</v-icon>
        <v-switch class="mt-5 ml-3 mr-n2" v-model="$vuetify.theme.dark" inset></v-switch>
        <v-icon>mdi-brightness-2</v-icon>
      </v-toolbar-items>
    </v-app-bar>

    <v-content>
      <v-container fluid class="fill-height d-flex flex-wrap">
        <container-tile
          v-for="c in containers"
          :container="c"
          :key="c.Hostname"
          v-bind:class="{'ma-1': true, half : containers.length > 1}"
        ></container-tile>
      </v-container>
    </v-content>
    <v-footer app fixed>
      <span style="margin-left:1em">&copy; You</span>
    </v-footer>
  </v-app>
</template>

<script>
import Wails from "@wailsapp/runtime";

import HelloWorld from "./components/HelloWorld.vue";
import Containers from "./components/Containers.vue";
import ContainerTile from "./components/ContainerTile.vue";

export default {
  data: () => ({
    containers: [],
    drawer: false
  }),

  async mounted() {
    window.backend.Containers
      .GetContainers()
      .then(res => {
        this.containers = res
        console.log('containers', this.containers)
      })
      .catch(error => {
        console.log("Error getting containers: ", error)
      })
  },

  components: {
    // HelloWorld,
    Containers,
    ContainerTile
  },
  methods: {}
};
</script>

<style>
.logo {
  width: 16em;
}
/* .container-tile {
  max-width: 33%;
} */
.tile-text {
  height: 315px;
  overflow: auto;
}
.half {
  width: 49%;
}
</style>