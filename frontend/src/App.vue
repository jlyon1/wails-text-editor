<template>
  <div ref="cont" id="app">
    <input type="text" v-model="filename"/>
    <button @click="save">
      SAVE
    </button>
    <select v-model="selected">
      <option v-for="option in options" :key="option">{{option}}</option>
    </select>
    <editor v-model="content" @init="editorInit" :lang="selected" theme="chrome" :width="editorWidth" :height="editorHeight"></editor>
    <div v-if="selected=='html'" style="margin-top: 30px; border: 1px solid black; min-height: 200px;" v-html="content"></div>
  </div>
</template>

<script>
import "./assets/css/main.css";
import editor from 'vue2-ace-editor'

export default {
  name: "app",
  components: {
    editor
  },
  computed:{
    editorWidth(){
      return 640
    },
    editorHeight(){
      return 480
    }
  },
  mounted(){
  },
  methods: {
    save(){
      window.backend.Save(this.filename,this.content)
    },
    editorInit: function () {
        require('brace/ext/language_tools') //language extension prerequsite...
        require('brace/mode/html')                
        require('brace/mode/javascript')    //language
        require('brace/mode/less')
        require('brace/theme/chrome')
        require('brace/snippets/javascript') //snippet
    }
  },
  data(){
    return {
      content: '',
      filename: '',
      selected: 'html',
      options: ['html', 'javascript']
      }
  }
};
</script>
