import React from 'react';

import {
    AlertDialog,
    AlertDialogBody,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogContent,
    AlertDialogOverlay,
    AlertDialogCloseButton,
    Button,
    useDisclosure,
    border
  } from '@chakra-ui/react';

import { useNavigate } from 'react-router-dom';


export default function TransitionExample() {
    const { isOpen, onOpen, onClose } = useDisclosure()
    const cancelRef = React.useRef()

    const navigate = useNavigate();
    const handleSubmit = (event) => {
        event.preventDefault(); // Evita el envío del formulario por defecto

        const formData = new FormData(event.target);

        //AQUI VAMOS A ESPECIFICAR LA RUTA A LA QUE VAMOS A MANDAR NUESTRO FORMULARIO
        fetch('/action_page.php', {
            method: 'POST',
            body: formData
        })
        .then(response => {
            if (response.ok) {
                // Redirigir a otra ruta después de un envío exitoso
                //navigate('/');
            } else {
                // Manejar errores
                navigate('/');
                console.error('Error en el envío del formulario');
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });
    };

    return (
      <>
        <Button onClick={onOpen} colorScheme='blue'  width='20%' >Continuar</Button>
        <AlertDialog
          motionPreset='slideInBottom'
          leastDestructiveRef={cancelRef}
          onClose={onClose}
          isOpen={isOpen}
          isCentered
        >
          <AlertDialogOverlay />
  
          <AlertDialogContent>
            <AlertDialogHeader>Permiso de superusuario</AlertDialogHeader>
            <AlertDialogCloseButton />

            <AlertDialogBody>
              Porfavor digite la contraseña de autorizacion 
              <form onSubmit={handleSubmit} method="post">
                <input className="" type="password" style={{border: '1px solid gray'}}/>
                <AlertDialogFooter>              
                        <Button colorScheme='blue' ml={3} display='flex' justifyContent='right' type='submit' >
                        Confirmar
                    </Button>
                </AlertDialogFooter>
              </form>
              
            </AlertDialogBody>

            
          </AlertDialogContent>
        </AlertDialog>
      </>
    )
  }