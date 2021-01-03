<script>
  import Card from "@smui/card";

  import { Col, Row } from "components/grid";
  import formatMoney from "utils/formatMoney";
  import formatDate from "utils/formatDate";

  export let cents = null;
  export let fireGoal = null;
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

  .fireGoal {
    margin-top: 1em;
    padding: 0 2em;
    text-align: center;
    font-weight: 600;
    font-size: 1.1em;
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
        {#if fireGoal}
          <div class="fireGoal">
            That's
            {((cents / fireGoal) * 100).toFixed(2)}% of your goal of
            {formatMoney(fireGoal)}
          </div>
        {/if}
        {#if lastUpdated}
          <div class="updatedTime">Last updated: {formatDate(lastUpdated)}</div>
        {/if}
      {/if}
    </Col>
  </Row>
</Card>
