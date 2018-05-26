package student_manager

import "testing"

func TestStudents_Add(t *testing.T) {
	stus := Students{}
	stu := Student{1, "test_stu_add"}
	stus.Add(stu)

	_, ok := stus[stu.Name]
	if !ok {
		t.Errorf("Students Add methods test fail: %v", stu)
	}

}

func TestStudents_Update(t *testing.T) {
	stus := Students{}
	stu := Student{1, "test_old_value"}
	stus.Add(stu)
	stus.Update("test_old_value", "test_new_value")

	_, ok := stus["test_new_value"]
	if !ok {
		t.Errorf("Students Updat methods test fail: %v", stu)

	}

}

func TestStudents_Delete(t *testing.T) {
	stus := Students{}
	stu := Student{1, "test_name"}
	stus.Add(stu)

	delete(stus, "test_name")

	_, exist := stus["test_name"]
	if exist {
		t.Errorf("Students Delete method test fail: %v", stu)
	}
}

func TestStudents_List(t *testing.T) {

}

func TestStudents_Load(t *testing.T) {

}

func TestStudents_Save(t *testing.T) {

}
