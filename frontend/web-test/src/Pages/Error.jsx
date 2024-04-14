import React from "react";
import { Container, Row, Col } from "react-bootstrap";

function ErrorPage({ errorCode }) {
  let errorMessage = "";


  switch (errorCode) {
    case 404:
      errorMessage = "Page Not Found";
      break;
    case 500:
      errorMessage = "Internal Server Error";
      break;
    default:
      errorMessage = "Unknown Error";
  }

  return (
    <Container>
      <Row className="mt-5">
        <Col>
          <h1>Error {errorCode}</h1>
          <p>{errorMessage}</p>
        </Col>
      </Row>
    </Container>
  );
}

export default ErrorPage;
