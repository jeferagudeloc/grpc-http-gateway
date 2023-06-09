
import { createTheme, ThemeProvider } from '@mui/material/styles';

import CssBaseline from '@mui/material/CssBaseline';
import { ProjectDescription } from './components/organism/ProjectDescription';
import { GlobalStyles } from '@mui/material';
import { VelocityCards } from './components/organism/VelocityCards';
import { requests } from './constants/requests';
import { Provider } from 'react-redux';
import store from "./store/store";

const themeLight = createTheme({
  palette: {
    background: {
      default: "#E8E9E8E9"
    }
  }
});

export default function App() {
  return (
    <Provider store={store}>
    <ThemeProvider theme={themeLight}>
      <GlobalStyles styles={{ ul: { margin: 0, padding: 0, listStyle: 'none' } }} />
      <CssBaseline />
      <ProjectDescription />
      <VelocityCards requests={requests} />
    </ThemeProvider>
    </Provider>
  );
}