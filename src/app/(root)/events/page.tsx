import { getEvents } from "@/api";
import { RouteGuard } from "@/components/auth";
import EventTable from "@/components/events/table/Table";

async function Page() {

  const events = await getEvents();

  return (
    <RouteGuard>
      <EventTable events={events} />
    </RouteGuard>
  )
}

export default Page
