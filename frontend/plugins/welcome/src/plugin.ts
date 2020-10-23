import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import CreateUser from './components/RecordUser';
import ShowUser from './components/RecordUserTable';
import CreateReturninvoice from './components/RecordReturninvoice';
import ShowReturninvoice from './components/RecordReturninvoiceTable';
import CreateRepairInvoice from './components/RecordRepairInvoice';
import ShowRepairInvoice from './components/RecordRepairInvoiceTable';
import CreateBill from './components/RecordBill';
import ShowBill from './components/RecordBillTable';
import CreatePartorder from './components/PartOrder';
import ShowPartorder from './components/PartOrderTable';
import SignIn from './components/SignIn';
 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/Home', WelcomePage);
    router.registerRoute('/RecordUser', CreateUser);
    router.registerRoute('/RecordUserTable', ShowUser);
    router.registerRoute('/RecordReturninvoice', CreateReturninvoice);
    router.registerRoute('/RecordReturninvoiceTable', ShowReturninvoice);
    router.registerRoute('/RecordRepairInvoice', CreateRepairInvoice);
    router.registerRoute('/RecordRepairInvoiceTable', ShowRepairInvoice);
    router.registerRoute('/RecordBill', CreateBill);
    router.registerRoute('/RecordBillTable', ShowBill);
    router.registerRoute('/PartOrder', CreatePartorder);
    router.registerRoute('/PartOrderTable', ShowPartorder);
    router.registerRoute('/', SignIn);
  },
});