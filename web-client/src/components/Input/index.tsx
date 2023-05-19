interface InputProps {
  type: string;
  placeholder: string;
  label: string;
}

export default function Input({ type, placeholder, label }: InputProps) {
  return (
    <div>
      <label htmlFor={"formId-" + placeholder}>
        {label}
        <input
          id={"formId-" + placeholder}
          type={type}
          placeholder={placeholder}
        />
      </label>
    </div>
  );
}
