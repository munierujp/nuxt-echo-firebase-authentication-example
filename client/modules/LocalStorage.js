class LocalStorage {
  constructor (key) {
    this.key = key
  }

  static of (key) {
    return new this(key)
  }

  get () {
    return localStorage.getItem(this.key)
  }

  set (value) {
    localStorage.setItem(this.key, value)
  }

  remove () {
    localStorage.removeItem(this.key)
  }
}

export default LocalStorage
