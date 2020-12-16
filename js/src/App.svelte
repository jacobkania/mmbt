<script>
  import { Router, Link, Route } from "svelte-navigator";
  import Home from "./pages/home/Home.svelte";
  import Overview from "./pages/overview/Overview.svelte";
  import TopNav from "./components/nav/TopNav.svelte";
</script>

<style>
  .main {
    height: 100vh;
    margin: 0 auto;
    background-color: #ccc;
  }
</style>

<Router>
  <div class="main" id="main">
    <TopNav />
    <div>
      <Route path="/">
        <Home />
      </Route>
      <Route path="overview" component={Overview} />
    </div>
  </div>
</Router>
