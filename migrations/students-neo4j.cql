// 1. Create constraints to ensure each entity has a unique id
CREATE CONSTRAINT specialization_id IF NOT EXISTS FOR (s:Specialization) REQUIRE s.id IS UNIQUE;
CREATE CONSTRAINT course_id IF NOT EXISTS FOR (c:Course) REQUIRE c.id IS UNIQUE;
CREATE CONSTRAINT class_id IF NOT EXISTS FOR (cl:Class) REQUIRE cl.id IS UNIQUE;
CREATE CONSTRAINT material_id IF NOT EXISTS FOR (m:Material) REQUIRE m.id IS UNIQUE;
CREATE CONSTRAINT schedule_id IF NOT EXISTS FOR (sc:Schedule) REQUIRE sc.id IS UNIQUE;
CREATE CONSTRAINT student_id IF NOT EXISTS FOR (st:Student) REQUIRE st.id IS UNIQUE;
CREATE CONSTRAINT group_id IF NOT EXISTS FOR (g:Group) REQUIRE g.id IS UNIQUE;

// 2. Create sample nodes (You can skip or customize this part as needed)
CREATE (spec:Specialization {id: "spec1"});
CREATE (course:Course {id: "course1"});
CREATE (class:Class {id: "class1"});
CREATE (mat:Material {id: "material1"});
CREATE (sch:Schedule {id: "schedule1"});
CREATE (stud:Student {id: "student1"});
CREATE (grp:Group {id: "group1"});

// 3. Create relationships between nodes
MATCH (spec:Specialization {id: "spec1"}), (course:Course {id: "course1"})
  CREATE (spec)-[:HAS_COURSE]->(course);

MATCH (course:Course {id: "course1"}), (spec:Specialization {id: "spec1"})
  CREATE (course)-[:IN_PROGRAM]->(spec);

MATCH (course:Course {id:"course1"}), (class:Class {id:"class1"})
  CREATE (course)-[:HAS_CLASS]->(class);

MATCH (class:Class {id:"class1"}), (mat:Material {id:"material1"})
  CREATE (class)-[:HAS_MATERIAL]->(mat);

MATCH (class:Class {id:"class1"}), (sch:Schedule {id:"schedule1"})
  CREATE (class)-[:TAKES_TIME_IN]->(sch);

MATCH (stud:Student {id:"student1"}), (sch:Schedule {id:"schedule1"})
  CREATE (stud)-[:ATTENDS]->(sch);

MATCH (grp:Group {id:"group1"}), (sch:Schedule {id:"schedule1"})
  CREATE (grp)-[:HAS_CLASS_IN]->(sch);

MATCH (grp:Group {id:"group1"}), (stud:Student {id:"student1"})
  CREATE (grp)-[:HAS_STUDENT]->(stud);

MATCH (stud:Student {id:"student1"}), (grp:Group {id:"group1"})
  CREATE (stud)-[:HAS_GROUP]->(grp);