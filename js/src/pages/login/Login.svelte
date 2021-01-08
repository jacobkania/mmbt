<script>
  import { Col, Row } from "components/grid";

  import TextField from "@smui/textfield";
  import Card, { Actions } from "@smui/card";
  import Button, { Label, Icon } from "@smui/button";
  import Snackbar, { Label as SBLabel } from "@smui/snackbar";

  let username = "";
  let password = "";
  let isSubmitting = false;
  let failureSnackbar;
  let failureSnackbarText;

  const handleSubmit = () => {
    isSubmitting = true;
    console.log("username: ", username);
    console.log("password: ", password);
    fetch("/login", {
      method: "POST",
      headers: new Headers({ "Content-Type": "application/json" }),
      body: JSON.stringify({
        username,
        passw: password,
      }),
    }).then((response) => {
      isSubmitting = false;
      console.log("response.ok: ", response.ok);
      response.json().then((body) => {
        console.log("BODY: ", body);
        if (!response.ok) {
          failureSnackbarText = body.error;
          failureSnackbar.open();
          return;
        }

        // set token cookie
        document.cookie = "token=" + JSON.stringify(body);
        window.location.reload();
      });
    });
  };
</script>

<style type="text/scss">
  @import "theme";

  .container {
    width: 100%;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .actions {
    margin-top: 1rem;
  }
</style>

<div class="container">
  <Card padded>
    <Row>
      <Col xs={12}>
        <TextField
          bind:value={username}
          disabled={isSubmitting}
          label="Username"
          fullwidth />
      </Col>
      <Col xs={12}>
        <TextField
          bind:value={password}
          disabled={isSubmitting}
          label="Password"
          fullwidth
          type="password" />
      </Col>
    </Row>
    <Actions fullBleed class="actions">
      <Button on:click={handleSubmit}>
        <Label>Log in</Label>
        <Icon class="material-icons">arrow_forward</Icon>
      </Button>
    </Actions>
  </Card>
</div>

<Snackbar bind:this={failureSnackbar}>
  <SBLabel>{failureSnackbarText}</SBLabel>
</Snackbar>
