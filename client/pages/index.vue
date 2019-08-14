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
    <hr>
    <div>
      <button
        :disabled="isLoading"
        @click="kickPublicApi"
      >
        Kick public API
      </button>
      <button
        :disabled="isLoading"
        @click="kickPrivateApi"
      >
        Kick private API
      </button>
    </div>
    <div>
      <span>Response: {{ response }}</span>
    </div>
  </div>
</template>

<script>
import firebase from '~/modules/firebase'

const { API_ORIGIN } = process.env.config
const PUBLIC_API_URL = `${API_ORIGIN}/public`
const PRIVATE_API_URL = `${API_ORIGIN}/private`

export default {
  asyncData () {
    return {
      isLoading: true,
      isSignedIn: false,
      user: {},
      response: ''
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
    },
    async kickPublicApi () {
      const response = await fetchText(PUBLIC_API_URL)
      this.response = response
    },
    async kickPrivateApi () {
      const response = await fetchText(PRIVATE_API_URL)
      this.response = response
    }
  }
}

async function fetchText (url) {
  const response = await fetch(url)
  const text = await response.text()
  return text
}
</script>
