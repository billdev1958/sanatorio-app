export interface servicios {
    id: number
    name: string
}

export interface pacientes{
    medical_history_id: string
    name: string
}

export interface appoitnment {
    dayOfWeek: number
    timeStart: string
    timeEnd: string
}