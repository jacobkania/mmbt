<script>
  import Card, { Content, Actions, ActionIcons } from "@smui/card";
  import IconButton, { Icon } from "@smui/icon-button";

  import { Col, Row } from "components/grid";
  import formatMoney from "utils/formatMoney";

  export let debts = [];
</script>

<style type="text/scss">
  @import "theme";

  .table {
    width: 100%;
    font-size: 1.2em;
  }

  .debtAmount {
    background-color: $money-outgo;
    color: $white;
    border-radius: 0.5em;
    padding: 0.5em;
  }
</style>

<Card>
  <Content>
    <Row>
      <Col>
        <h2>Debts:</h2>
        <table class="table">
          {#each debts as item}
            <tr>
              <td>{item.name}</td>
              <td class="debtAmount">{formatMoney(item.cents)}</td>
            </tr>
          {:else}
            <div>There's nothing here</div>
          {/each}
        </table>
      </Col>
    </Row>
  </Content>
  <Actions>
    <ActionIcons>
      <IconButton class="material-icons" on:click={() => alert('click')}>
        edit
      </IconButton>
    </ActionIcons>
  </Actions>
</Card>
