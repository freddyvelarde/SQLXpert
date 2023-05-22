import { useParams } from "react-router-dom";

export default function Dinamic() {
  const { id } = useParams();
  return <h1>dinamix page: {id}</h1>;
}
