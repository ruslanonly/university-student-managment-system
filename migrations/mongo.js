db = db.getSiblingDB('my_database');

db.universities.insertMany([
  {
    name: "Московский институт радиотехники, электроники и автоматики (МИРЭА)",
    institutes: [
      {
        name: "Институт информационных технологий",
        departments: [
          { name: "Кафедра программного обеспечения" },
          { name: "Кафедра радиотехники" }
        ]
      },
      {
        name: "Институт искусственного интеллекта",
        departments: [
          { name: "Кафедра автоматического управления" }
        ]
      },
      {
        name: "Институт радиоэлектроники и информатики",
        departments: []
      },
      {
        name: "Институт кибербезопасноти и цифровых технологий",
        departments: []
      },
      {
        name: "Институт технологий управления",
        departments: []
      },
      {
        name: "Институт тонких химических технологий им. М.В. Ломоносова",
        departments: []
      },
      {
        name: "Институт перспективных технологий и индустриального программирования",
        departments: []
      }
    ]
  }
]);

db.groups.insertMany([
  {
    group_id: 1,
    group_name: "БСБО-01-21",
    students: [
      { student_id: 1, full_name: "Петров Алексей Сергеевич" },
      { student_id: 2, full_name: "Смирнов Дмитрий Валерьевич" },
      { student_id: 3, full_name: "Кузнецова Мария Александровна" },
      { student_id: 4, full_name: "Васильев Николай Игоревич" },
      { student_id: 5, full_name: "Соколова Ольга Вячеславовна" },
      { student_id: 6, full_name: "Григорьев Андрей Викторович" },
      { student_id: 7, full_name: "Зайцева Екатерина Анатольевна" },
      { student_id: 8, full_name: "Михайлов Артем Сергеевич" },
      { student_id: 9, full_name: "Попова Валерия Евгеньевна" },
      { student_id: 10, full_name: "Кравцов Кирилл Александрович" },
      { student_id: 11, full_name: "Иванов Иван Иванович" }
    ]
  },
  {
    group_id: 2,
    group_name: "БСБО-02-21",
    students: [
      { student_id: 12, full_name: "Фадеев Всеволод Вадимович" },
      { student_id: 13, full_name: "Усанкин Александр Александрович" },
      { student_id: 14, full_name: "Рзаев Руслан Халидович" }
    ]
  },
  {
    group_id: 3,
    group_name: "БСБО-04-21",
    students: [
      { student_id: 15, full_name: "Тимофеев Максим Павлович" },
      { student_id: 16, full_name: "Лебедев Артем Юрьевич" },
      { student_id: 17, full_name: "Фролова Дарина Васильевна" },
      { student_id: 18, full_name: "Ильин Иван Игоревич" },
      { student_id: 19, full_name: "Чернова Светлана Павловна" },
      { student_id: 20, full_name: "Андреев Павел Владиславович" },
      { student_id: 21, full_name: "Егорова Виктория Николаевна" },
      { student_id: 22, full_name: "Морозов Михаил Евгеньевич" },
      { student_id: 23, full_name: "Сидорова Ирина Борисовна" },
      { student_id: 24, full_name: "Борисов Сергей Юрьевич" }
    ]
  }
]);
