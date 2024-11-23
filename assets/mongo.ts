type University = {
  id: number
  name: string
  institutes: Institute[]
}

type Institute = {
  id: number
  name: string
  departments: Department[]
}

type Department = {
  id: number
  name: string
}
