import { useParams } from "react-router-dom";

export default function Dashboard() {
  const { database } = useParams();
  return <h1>Dashboard: {database}</h1>;
}
