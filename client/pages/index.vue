<template>
  <div v-if="isLoading">
    <span>Now loading...</span>
  </div>
  <div v-else>
    <div v-if="isSignedIn">
      <div>
        <span>You are signed in as {{ user.displayName }} ({{ user.email }}).</span>
      </div>
      <div>
        <button
          :disabled="isLoading"
          @click="signOut"
        >
          Sign out
        </button>
      </div>
    </div>
    <div v-else>
      <div>
        <span>You are not signed in.</span>
      </div>
      <div>
        <button
          :disabled="isLoading"
          @click="signIn"
        >
          Sign in with Google
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import firebase from '~/modules/firebase'

export default {
  asyncData () {
    return {
      isLoading: true,
      isSignedIn: false,
      user: {}
    }
  },
  mounted () {
    firebase.auth().onAuthStateChanged((user) => {
      this.isLoading = false
      if (user) {
        this.isSignedIn = true
        this.user = user
      } else {
        this.isSignedIn = false
        this.user = {}
      }
    })
  },
  methods: {
    signIn () {
      const provider = new firebase.auth.GoogleAuthProvider()
      firebase.auth().signInWithRedirect(provider)
    },
    signOut () {
      firebase.auth().signOut()
    }
  }
}
</script>
