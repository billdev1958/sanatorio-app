import * as React from 'react'
import * as ReactDOM from 'react-dom/client'
import { ChakraProvider } from '@chakra-ui/react'     

import PaginaPrincipal from './routes/PaginaPrincipal'
import PaginaRegistro from './routes/PaginaRegistro';
import PRSUser from './routes/PRSUser';

import ErrorPage from "./ErrorPage";
import Contact from "./routes/Contact";



import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
const router = createBrowserRouter([
  {
    path: "/",
    element: <PaginaPrincipal />,
    errorElement: <ErrorPage />,
  },
  {
    path: "/registro",
    element: <PaginaRegistro />,
    errorElement: <ErrorPage />,
    children: [
      {
        path: "/registro/admin",
        element: <PRSUser />,
      },
    ],
  },
  {
    path: "/registro/1",
    element: <PaginaRegistro />,
    errorElement: <ErrorPage />,
  },
  {
    path: "/registro/2",
    element: <PaginaRegistro />,
    errorElement: <ErrorPage />,
  },
  {
    path: "/registro/3",
    element: <PaginaRegistro />,
    errorElement: <ErrorPage />,
  },
  
]);



ReactDOM.createRoot(document.getElementById("root")).render(
  <ChakraProvider>
    <React.StrictMode>
      
        <RouterProvider router={router} />
      
    </React.StrictMode>
  </ChakraProvider>
);