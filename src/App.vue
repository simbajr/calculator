<template>
  <!-- <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Welcome to Your Vue.js App"/> -->
  <div class="hello">

    <div>
      <h1 class="title"> The Calculator </h1>
    </div>

    <hr>
    <div class="container">
      <div>
        <div class="col">
          <form>

            <div class="col-sm-12">
              <label class="col-sm-2"> First number </label>
              <input type="text" class="form-group col-sm-2" v-model="num1" placeholder="Num 1">
            </div>

            <div style="margin-top:13px" class="col-sm-12">
              <label class="col-sm-2"> Second number </label>
              <input type="text" class="form-group col-sm-2" v-model="num2" placeholder="Num 2">

            </div>

          </form>
        </div>
        <div class="col">
          <div class="col-sm-12">
            <div><label class="col-sm-2">Addition: {{ add }}</label></div>
            <div><label class="col-sm-2">Multiplication: {{ mul }}</label></div>
            <div><label class="col-sm-2">Subtraction: {{ sub }}</label></div>
            <div><label class="col-sm-2">Division: {{ div }}</label></div>
          </div>
        </div>
      </div>

      <div class="col">
        <div class="col-sm-4">
          <button type="button" class="btn btn-primary" v-on:click="postRequest()">Calcuclate</button>
        </div>
      </div>
    </div>
    &nbsp;

    &nbsp;
  </div>
</template>

<script>

import axios from 'axios'


export default {
  name: 'App',
  data: function () {
    return {
      add: "", mul: "", sub: "", div: "",
      num1: "", num2: ""
    }
  },
  methods: {
    postRequest() {
      var data = { "num1": parseFloat(this.num1), "num2": parseFloat(this.num2) }


      console.log(data)

      axios({ method: "POST", url: "http://127.0.0.1:8088/calc", data: data, headers: { "content-type": "text/plain" } }).then(result => {
        //caseSensative
        this.add = result.data['Add']
        this.mul = result.data['Mul']
        this.sub = result.data['Sub']
        this.div = result.data['Div']

        console.log(result.data)
      }).catch(error => {
        console.log(error)
      })
    }
  }
}
</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
