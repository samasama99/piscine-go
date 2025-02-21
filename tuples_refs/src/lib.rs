pub struct Student(pub u32, pub String, pub String);

pub fn id(student: &Student) -> u32 {
    student.0
}

pub fn first_name(student: &Student) -> &str {
    student.1.as_str()
}

pub fn last_name(student: &Student) -> &str {
    student.2.as_str()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let student = Student(20, "Pedro".to_string(), "Domingos".to_string());
        assert_eq!(20, id(&student));
        assert_eq!("Pedro", first_name(&student));
        assert_eq!("Domingos", last_name(&student));
    }
}
