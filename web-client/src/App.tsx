import styled from "styled-components";
import { testUrl } from "./data/endpoints";

const TestStyles = styled.h1`
  color: orange;
`;

function App() {
  async function getData() {
    const res = await fetch(testUrl);
    const data = await res.json();
    console.log(data);
  }

  return (
    <>
      <TestStyles>Hello world</TestStyles>
      <button onClick={getData}>get data</button>
    </>
  );
}

export default App;
