<script>
  import Card, { Content, Actions, ActionIcons } from "@smui/card";
  import IconButton, { Icon } from "@smui/icon-button";

  import { Col, Row } from "components/grid";
  import formatMoney from "utils/formatMoney";

  export let budgetItems = [];
</script>

<style type="text/scss">
  @import "theme";

  .table {
    width: 100%;
    font-size: 1.2em;
  }

  .budgetAmount {
    background-color: $tertiary;
    color: $white;
    border-radius: 0.5em;
    padding: 0.5em;
  }
</style>

<Card>
  <Content>
    <Row>
      <Col>
        <h2>Monthly Budget:</h2>
        <table class="table">
          {#each budgetItems as item}
            <tr>
              <td>{item.name}</td>
              <td class="budgetAmount">{formatMoney(item.cents)}</td>
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
