<script>
  import Card from "@smui/card";

  import { Col, Row } from "components/grid";
  import formatMoney from "utils/formatMoney";
  import formatDate from "utils/formatDate";

  export let cents = null;
  export let lastUpdated = null;
</script>

<style type="text/scss">
  @import "theme";

  .netWorthContainer {
    background-color: $secondary;
    color: $white;
    border-radius: 2em;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .netWorthVal {
    display: flex;
    font-size: 2em;
    padding: 0.5em;
  }

  .updatedTime {
    margin-top: 1em;
  }
</style>

<Card padded>
  <Row>
    <Col>
      <h2>Current net worth:</h2>
      {#if cents}
        <div class="netWorthContainer">
          <span class="netWorthVal">{formatMoney(cents)}</span>
        </div>
        {#if lastUpdated}
          <div class="updatedTime">Last updated: {formatDate(lastUpdated)}</div>
        {/if}
      {/if}
    </Col>
  </Row>
</Card>
