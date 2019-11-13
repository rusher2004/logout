<template>
  <v-container>
    <v-layout>
      <v-flex xs12 sm6 offset-sm3>
        
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  props: {
    details: Object
  },
  created() {
    console.log('props', this.props)
  }
}
</script>