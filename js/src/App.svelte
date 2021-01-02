<script>
  import { Route } from "tinro";
  import Home from "pages/home/Home.svelte";
  import Overview from "pages/overview/Overview.svelte";
  import TopNav from "components/nav/TopNav.svelte";
</script>

<style type="text/scss">
  @import "theme";

  .main {
    height: 100vh;
    margin: 0 auto;
    color: $text;
  }
</style>

<Route>
  <div class="main" id="main">
    <TopNav />
    <div>
      <Route path="/">
        <Home />
      </Route>
      <Route path="/overview">
        <Overview />
      </Route>
      <Route fallback>NOTHIN</Route>
    </div>
  </div>
</Route>
