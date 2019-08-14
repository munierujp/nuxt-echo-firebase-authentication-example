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
import LocalStorage from '~/modules/LocalStorage'

const JWT = LocalStorage.of('jwt')

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
      firebase.auth().signInWithPopup(provider).then((res) => {
        res.user.getIdToken().then((idToken) => {
          JWT.set(idToken.toString())
        })
      })
    },
    signOut () {
      firebase.auth().signOut().then(() => {
        JWT.remove()
      })
    },
    async kickPublicApi () {
      const response = await fetch(PUBLIC_API_URL)
      const text = await response.text()
      this.response = text
    },
    async kickPrivateApi () {
      const jwt = JWT.get()
      const response = await fetch(PRIVATE_API_URL, {
        headers: {
          'Authorization': `Bearer ${jwt}`
        }
      })
      const text = await response.text()
      this.response = text
    }
  }
}
</script>
