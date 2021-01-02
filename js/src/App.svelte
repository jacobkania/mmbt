<script>
  import { Route } from "tinro";

  import TopNav from "components/nav/TopNav.svelte";

  import Home from "pages/home/Home.svelte";
  import Overview from "pages/overview/Overview.svelte";
  import NewSnapshot from "pages/snapshot/new/NewSnapshot.svelte";
</script>

<style type="text/scss">
  @import "theme";

  .main {
    height: 100vh;
    margin: 0 auto;
    color: $text;
    font-family: "fira", sans-serif;
  }
</style>

<Route>
  <div class="main" id="main">
    <TopNav />
    <div id="content">
      <Route path="/">
        <Home />
      </Route>
      <Route path="/overview">
        <Overview />
      </Route>
      <Route path="/snapshot/*">
        <Route path="/new">
          <NewSnapshot />
        </Route>
      </Route>
      <Route fallback>Page not found</Route>
    </div>
  </div>
</Route>
