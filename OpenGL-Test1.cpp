#include <GL/glut.h>

void desenho()
{
	glClear(GL_COLOR_BUFFER_BIT);
	glLineWidth(5);

	// fazendo a casa 1 estrutura

	glBegin(GL_QUADS);
	glColor3f(0, 0.5, 0.5);
	glVertex2f(-3, 3);
	glVertex2f(-3, -3);
	glVertex2f(3, -3);
	glVertex2f(3, 3);
	glEnd();

	// fazendo a casa 2 telhado

	glBegin(GL_TRIANGLES);
	glColor3f(0, 0.5, 1);
	glVertex2f(0, 6);
	glVertex2f(-3, 3);
	glVertex2f(3, 3);
	glEnd();

	// fazendo a casa 3 porta

	glColor3f(0, 1, 0.5);
	glBegin(GL_QUADS);
	glVertex2f(-1, -1);
	glVertex2f(-1, -3);
	glVertex2f(0, -3);
	glVertex2f(0, -1);
	glEnd();

	// fazendo a casa 4 fechadura

	glColor3f(0, 0, 0);
	

	glFlush();
}

int main(int argc, char** argv)
{
	glutInit(&argc, argv);
	glutInitDisplayMode(GLUT_SINGLE | GLUT_RGBA);
	glutInitWindowSize(600, 600);
	glutInitWindowPosition(10, 20);
	glutCreateWindow("Primitivas");
	glutDisplayFunc(desenho);
	gluOrtho2D(-10, 10, -10, 10);
	glClearColor(256.0 / 256.0, 256.0 / 256.0, 256.0 / 256.0, 0);
	glutMainLoop();
	return 0;
}

