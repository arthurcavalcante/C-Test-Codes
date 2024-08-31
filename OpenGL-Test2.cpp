#include <GL/glut.h>

void display()
{
	glClear(GL_COLOR_BUFFER_BIT);

	glColor3f(153, 124, 61);
	glBegin(GL_QUADS);
	glVertex2f(200, 400);
	glVertex2f(200, 100);
	glVertex2f(600, 100);
	glVertex2f(600, 400);
	glEnd();

	glColor3f(173, 138, 61);
	glBegin(GL_TRIANGLES);
	glVertex2f(400, 600);
	glVertex2f(200, 400);
	glVertex2f(600, 400);
	glEnd();

	glColor3f(0, 0, 0);
	glBegin(GL_QUADS);
	glVertex2f(300, 200);
	glVertex2f(300, 100);
	glVertex2f(400, 100);
	glVertex2f(400, 200);
	glEnd();

	glColor3f(0, 0, 0.8);
	glBegin(GL_QUADS);
	glVertex2f(400, 400);
	glVertex2f(400, 300);
	glVertex2f(500, 300);
	glVertex2f(500, 400);
	glEnd();

	
	glFlush();
	glutSwapBuffers();
}

int main(int argc, char** argv)
{
	glutInit(&argc, argv);
	glutInitDisplayMode(GLUT_DOUBLE | GLUT_RGBA);
	glutInitWindowPosition(0, 0);
	glutInitWindowSize(800, 800);
	glutCreateWindow("OpenGL");
	glutDisplayFunc(display);
	gluOrtho2D(0, 640, 0, 640);
	glClearColor(0, 0.1, 0, 1);
	glutMainLoop();
	return 0;
}
