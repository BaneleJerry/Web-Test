import React from "react";
import { Container, Row, Col, Button } from "react-bootstrap";

function Home() {
  return (
    <Container>
      <Row className="mt-5">
        <Col>
          <h1>Welcome to My Jogon Home</h1>
          <p>Where everything is a little bit quirky!</p>
        </Col>
      </Row>
      <Row className="mt-4">
        <Col>
          <h2>Jogon Features:</h2>
          <ul>
            <li>Flashing neon lights</li>
            <li>Giant rubber duck pond</li>
            <li>Gravity-defying furniture</li>
            <li>Cupcake dispenser</li>
          </ul>
        </Col>
      </Row>
      <Row className="mt-4">
        <Col>
          <Button variant="warning">Experience Jogon Now!</Button>
        </Col>
      </Row>
      <Row className="mt-5">
        <Col>
          <img
            src="https://via.placeholder.com/600"
            alt="Jogon House"
            className="img-fluid"
          />
        </Col>
      </Row>
    </Container>
  );
}

export default Home;
