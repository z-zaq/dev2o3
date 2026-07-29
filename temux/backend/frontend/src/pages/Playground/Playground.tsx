import Button from "../../components/ui/Button/Button";

export default function Playground() {
  return (
    <main className="min-h-screen bg-slate-100 p-12">
      <div className="mx-auto max-w-7xl space-y-12">
        <section>
          <h1 className="mb-8 text-4xl font-bold">
            Temux UI Playground
          </h1>

          <div className="flex flex-wrap gap-4">
            <Button>Primary</Button>

            <Button variant="secondary">
              Secondary
            </Button>

            <Button variant="outline">
              Outline
            </Button>

            <Button variant="ghost">
              Ghost
            </Button>
          </div>
        </section>
      </div>
    </main>
  );
}