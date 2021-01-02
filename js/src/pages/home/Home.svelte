<script>
  import { Col, Row } from "components/grid";

  import NetWorth from "./components/NetWorth.svelte";
  import Income from "./components/Income.svelte";
  import MonthlyBudget from "./components/MonthlyBudget.svelte";
  import SpendingGoals from "./components/SpendingGoals.svelte";

  let clicked = 0;
</script>

<style type="text/scss">
</style>

<svelte:head>
  <title>FI Budget | Home</title>
</svelte:head>

<Row>
  <Col xs="12" sm="6">
    <NetWorth cents="12340000" lastUpdated="2020-01-17T10:23:01Z" />
  </Col>
  <Col xs="12" sm="6">
    <Income
      incomes={[{ name: 'Burger Restaurant', cents: 90000 }, { name: 'Carwash', cents: 14000 }]} />
  </Col>
  <Col xs="12" sm="6">
    <MonthlyBudget
      budgetItems={[{ name: 'Living expenses', cents: 70000 }, { name: 'Car', cents: 40000 }]} />
  </Col>
</Row>
