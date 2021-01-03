<script>
  import { Col, Row } from "components/grid";

  import NetWorth from "./components/NetWorth.svelte";
  import Income from "./components/Income.svelte";
  import AccountValues from "./components/AccountValues.svelte";
  import Assets from "./components/Assets.svelte";
  import Expenses from "./components/Expenses.svelte";
  import Debts from "./components/Debts.svelte";

  let clicked = 0;
</script>

<style type="text/scss">
</style>

<svelte:head>
  <title>FI Budget | Home</title>
</svelte:head>

<Row>
  <Col xs="12" sm="6">
    <NetWorth
      cents={12340000}
      fireGoal={300000000}
      lastUpdated="2020-01-17T10:23:01Z" />
  </Col>
  <Col xs="12" sm="6">
    <Income
      incomes={[{ name: 'Burger Restaurant', cents: 90000 }, { name: 'Carwash', cents: 14000 }]} />
  </Col>
  <Col xs="12" sm="6">
    <AccountValues
      accounts={[{ name: 'Burger Restaurant', cents: 90000 }, { name: 'Carwash', cents: 14000 }]} />
  </Col>
  <Col xs="12" sm="6">
    <Assets
      assets={[{ name: 'Burger Restaurant', cents: 90000 }, { name: 'Carwash', cents: 14000 }]} />
  </Col>
  <Col xs="12" sm="6">
    <Expenses
      expenses={[{ name: 'Living expenses', cents: 70000 }, { name: 'Car', cents: 40000 }]} />
  </Col>
  <Col xs="12" sm="6">
    <Debts
      debts={[{ name: 'Living expenses', cents: 70000 }, { name: 'Car', cents: 40000 }]} />
  </Col>
</Row>
