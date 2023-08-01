export function Heading1({ children }: { children: React.ReactNode }) {
  return (
    <h1 className="font-heading text-2xl uppercase text-blue-50 2xl:text-4xl">
      {children}
    </h1>
  );
}
