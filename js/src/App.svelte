<script>
  import { Route } from "tinro";

  import Home from "pages/home/Home.svelte";
  import Overview from "pages/overview/Overview.svelte";
  import NewSnapshot from "pages/snapshot/new/NewSnapshot.svelte";
  import Login from "pages/login/Login.svelte";

  import TopNav from "components/nav/TopNav.svelte";

  import { token } from "stores/tokenStore";
</script>

<style type="text/scss">
  @import "theme";

  .main {
    min-height: 100vh;
    margin: 0 auto;
    color: $text;
    background-color: $background;
  }

  .content {
    margin: 0 $margin;
    max-width: 1920px;
  }
</style>

<svelte:head>
  <link rel="stylesheet" href="" type="text/scss" />
</svelte:head>

{#if $token}
  <Route>
    <div class="main" id="main">
      <TopNav />
      <div class="content" id="content">
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
{:else}
  <Route>
    <Login />
  </Route>
{/if}
