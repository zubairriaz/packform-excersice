<template>
    <header class="header">
    </header>
    <div class="container" id="app">
        <router-view></router-view>
    </div>
</template>

<script>
    export default {
        name: 'app'
    }
</script>

<style>
    body {
        background: #fff;
        color: #CACACC;
        font-family: 'Roboto Mono', monospace;
        font-size: 1em;
        justify-content: center;
        align-content: space-between;
        margin: 0px;
    }
    .container {
        margin: 0 auto;
        width: 80%;
    }
    h1 {
        font-weight: 400;
        margin: 0;
    }
    .header{
      width: 100%;
      height: 4em;
      background-color: #1d1d1d;
      margin-bottom: 20px;
    }
</style>