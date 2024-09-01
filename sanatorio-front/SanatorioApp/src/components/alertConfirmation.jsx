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

export default function TransitionExample() {

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
            <AlertDialogHeader>Registro exitoso</AlertDialogHeader>
            <AlertDialogCloseButton />

            <AlertDialogBody>
                            
            </AlertDialogBody>

            <AlertDialogFooter>              
                <Button colorScheme='blue' ml={3} display='flex' justifyContent='right' type='submit' >
                    Continuar
                </Button>
            </AlertDialogFooter>
            
          </AlertDialogContent>
        </AlertDialog>
      </>
    )
  }