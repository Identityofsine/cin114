import { RouteGuard } from '@/components/auth';
import DebugProvider from '@/components/auth/DebugProvider';
import './styles/dashboard.scss';

function DashboardContent() {
  return (
    <>
      <DebugProvider />
      <div className="dashboard">
        <div className="dashboard__content">
          <h1>Protected Dashboard</h1>
          <p>This page is only accessible to authenticated users.</p>
          <p>If you can see this, you are successfully logged in!</p>
        </div>
      </div>
    </>
  );
}

export default function DashboardPage() {
  return (
    <RouteGuard
      fallback={<div>Checking authentication...</div>}
      redirectTo="/zcpkcrucpw"
    >
      <DashboardContent />
    </RouteGuard>
  );
} 