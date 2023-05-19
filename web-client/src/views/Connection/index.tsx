import Input from "../../components/Input";

export default function Connection() {
  return (
    <div>
      <h1>connection view</h1>
      <form action="">
        <Input type="text" placeholder="localhost" label="HOST: " />
        <Input type="text" placeholder="5432" label="PORT: " />
        <Input type="text" placeholder="postgres" label="DB NAME: " />
        <Input type="text" placeholder="admin" label="DB PASSWORD: " />
      </form>
    </div>
  );
}
