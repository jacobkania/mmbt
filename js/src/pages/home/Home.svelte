<script>
  import Card from "@smui/card";
  import Button from "@smui/button";

  import { Col, Row } from "components/grid";
  import NetWorth from "./components/NetWorth.svelte";

  let clicked = 0;
</script>

<style type="text/scss">
</style>

<Row>
  <Col xs="12" sm="6">
    <NetWorth cents="12340000" lastUpdated="2020-01-17T10:23:01Z" />
  </Col>
  <Col xs="12" sm="6">
    Hello! this will have lots of words to make into a real column. a a a a a
    Hello! this will have lots of words to make into a real column. Hello! this
    will have lots of words to make into a real column.
  </Col>
</Row>
